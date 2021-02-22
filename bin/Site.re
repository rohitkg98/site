open Lib.Server;
open Lwt.Infix;

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

