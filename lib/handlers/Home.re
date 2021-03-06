open Tyxml;

let mycontent =
  <div className="content">
    <h1> "A fabulous title" </h1>
    "This is a fabulous content."
  </div>;

let mytitle = Html.txt("A Fabulous Web Page");

let mypage =
  <html>
    <head> <title> mytitle </title> </head>
    <body> mycontent </body>
  </html>;

let body = () => Format.asprintf("%a@.", Html.pp(~indent=true, ()), mypage);
