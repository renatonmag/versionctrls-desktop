/* @refresh reload */
import { render } from "solid-js/web";

import "./index.css";
import App from "./App";
import { A, HashRouter } from "@solidjs/router";

const root = document.getElementById("root");

if (import.meta.env.DEV && !(root instanceof HTMLElement)) {
  throw new Error(
    "Root element not found. Did you forget to add it to your index.html? Or maybe the id attribute got misspelled?"
  );
}

const Root = (props) => (
  <>
    <A href="/">Home</A>
    <A href="/hello-world">Hello World</A>
    <A href="/about">About</A>
    {props.children}
  </>
);

const routes = [
  {
    path: "/",
    component: () => <div>Repositories</div>,
  },
  {
    path: "/hello-world",
    component: () => <h1>Hello, World!</h1>,
  },
  {
    path: "/about",
    component: () => <div>Repository</div>,
  },
];

render(() => <HashRouter root={Root}>{routes}</HashRouter>, root!);
