let invalid_request: (Httpaf.Reqd.t, Httpaf.Status.t, string) => unit;
let request_handler: ('a, Httpaf.Reqd.t) => unit;
let error_handler:
  ('a, ~request: 'b=?, 'c, Httpaf.Headers.t => Httpaf.Body.t([ | `write])) =>
  unit;
