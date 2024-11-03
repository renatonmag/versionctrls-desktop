import { onMount } from "solid-js";
import { Button } from "./components/ui/button";
import { useStore } from "./store/store";
import { RepoListCards } from "./components/initial/RepoListCards";
import { ListRepos } from "wailsjs/go/backend/App";

export default function App() {
  const [_, { startWatching, openRepository, loadRepos }] = useStore();

  onMount(() => {
    loadRepos();
  });

  return (
    <div class="flex justify-center h-screen space-x-2">
      <div class="">
        <RepoListCards />
      </div>
      <div class="flex flex-col items-center space-y-2">
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
      </div>
    </div>
  );
}
