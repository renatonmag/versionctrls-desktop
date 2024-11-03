/* @refresh reload */
import { render } from "solid-js/web";

import "./index.css";
import App from "./App";
import { HashRouter } from "@solidjs/router";

render(
  () => <HashRouter root={App}>{/*... routes */}</HashRouter>,
  document.getElementById("root") as HTMLElement
);
