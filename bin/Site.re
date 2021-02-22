open Httpaf;
open Lwt.Infix;

let invalid_request = (reqd, status, body) => {
  /* Responses without an explicit length or transfer-encoding are
     close-delimited. */
  let headers = Headers.of_list([("Connection", "close")]);
  Reqd.respond_with_string(reqd, Response.create(~headers, status), body);
};

let request_handler = (_, reqd) => {
  let {Request.meth, target, _} = Reqd.request(reqd);
  switch (meth) {
  | `GET =>
    switch (String.split_on_char('/', target)) {
    | ["", "hello", ...rest] =>
      let who =
        switch (rest) {
        | [] => "world"
        | [who, ..._] => who
        };

      let response_body = Printf.sprintf("Hello, %s!\n", who);
      /* Specify the length of the response. */
      let headers =
        Headers.of_list([
          ("Content-length", string_of_int(String.length(response_body))),
        ]);

      Reqd.respond_with_string(
        reqd,
        Response.create(~headers, `OK),
        response_body,
      );
    | _ =>
      let response_body = Printf.sprintf("%S not found\n", target);
      invalid_request(reqd, `Not_found, response_body);
    }
  | meth =>
    let response_body =
      Printf.sprintf(
        "%s is not an allowed method\n",
        Method.to_string(meth),
      );

    invalid_request(reqd, `Method_not_allowed, response_body);
  };
};

let error_handler = (_client_address, ~request as _=?, error, start_response) => {
  /* We start the error response by calling the `start_response` function. We
   * get back a response body. */
  let response_body = start_response(Headers.empty);
  /* Once we get the response body, we can immediately start writing to it. In
   * this case, it might be sufficient to say that there was an error. */
  Body.write_string(
    response_body,
    "There was an error handling your request.\n",
  );
  /* Finally, we close the streaming response body to signal to the underlying
   * HTTP/2 framing layer that we have finished sending the response. */
  Body.close_writer(response_body);
};

let () = {
  /* We're going to be using the `H2_lwt_unix` module from the `h2-lwt-unix`
   * library to create a server that communicates over the underlying operating
   * system socket abstraction. The first step is to create a connection
   * handler that will accept incoming connections and let our request and
   * error handlers handle the request / response exchanges in those
   * connections. */
  let connection_handler =
    Httpaf_lwt_unix.Server.create_connection_handler(
      ~config=?None,
      ~request_handler,
      ~error_handler,
    );

  /* We'll be listening for requests on the loopback interface (localhost), on
   * port 8080. */
  let listen_address =
    Unix.([@implicit_arity] ADDR_INET(inet_addr_loopback, 8080));
  /* The final step is to start a server that will set up all the low-level
   * networking communication for us, and let it run forever. */

  let server = () => {
    Lwt_io.establish_server_with_client_socket(
      listen_address,
      connection_handler,
    )
    >|= ignore;
  };

  Lwt.async(server);
  let (forever, _) = Lwt.wait();
  Lwt_main.run(forever);
};

