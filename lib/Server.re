open Httpaf;
open Handlers;

let invalid_request = (reqd, status, body) => {
  /* Responses without an explicit length or transfer-encoding are
     close-delimited. */
  let headers = Headers.of_list([("Connection", "close")]);
  Reqd.respond_with_string(reqd, Response.create(~headers, status), body);
};

// attaches content length
let respond = (~headers=Headers.empty, body, reqd) => {
  let headers =
    Headers.add(
      headers,
      "Content-length",
      string_of_int(String.length(body)),
    );

  Reqd.respond_with_string(reqd, Response.create(~headers, `OK), body);
};

let request_handler = (_, reqd) => {
  let {Request.meth, target, _} = Reqd.request(reqd);
  switch (meth) {
  | `GET =>
    switch (String.split_on_char('/', target)) {
    | ["", "home"] => respond(Home.body(), reqd);
    | ["", "hello", ...rest] =>
      let who =
        switch (rest) {
        | [] => "world"
        | [who, ..._] => who
        };

      let response_body = Printf.sprintf("Hello, %s!\n", who);

      respond(response_body, reqd);
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
