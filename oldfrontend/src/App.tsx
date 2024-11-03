import { onMount, Suspense } from "solid-js";
import { Button } from "./components/ui/button";
import { StoreProvider, useStore } from "./store/store";
import { RepoListCards } from "./components/initial/RepoListCards";
import { ListRepos } from "wailsjs/go/backend/App";
import { MemoryRouter, createMemoryHistory, Route } from "@solidjs/router";

export default function App() {
  // const [_, { loadRepos }] = useStore();

  // onMount(() => {
  //   loadRepos();
  // });

  return <Repositories />;
}

const Root = (props) => (
  <>
    <h1>Root header</h1>
  </>
);

const Repositories = () => {
  return (
    <div class="flex justify-center items-center h-screen space-x-2">
      <div class="">
        <RepoListCards />
      </div>
      {/* <div class="flex flex-col items-center space-y-2">
        <Button class="w-fit" onClick={openRepository}>
          Open Repository
        </Button>
        <Button class="w-fit" onClick={startWatching}>
          Watch
        </Button>
        <Button
          class="w-fit"
          onClick={() => {
            ListRepos().then((repos) => {
              console.log({ repos });
            });
          }}
        >
          Verify Integration
        </Button>
      </div> */}
    </div>
  );
};

const Repository = () => {
  return <div>Repository</div>;
};
