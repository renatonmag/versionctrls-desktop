import { createEffect, createSignal, For } from "solid-js";

import { Button } from "../ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "../ui/card";
import { Switch, SwitchControl, SwitchThumb } from "../ui/switch";
import { Bell, Check, FolderUp, X } from "lucide-solid";
import { useStore } from "~/store/store";

const notifications = [
  {
    name: "versionctrls",
    path: "/Users/josh/code/versionctrls",
  },
  {
    name: "frames",
    path: "/Users/josh/code/frames",
  },
  {
    name: "solid-js",
    path: "/Users/josh/code/solid-js",
  },
];

export function RepoListCards() {
  const [state, { openRepository }] = useStore();
  return (
    <Card class="w-[380px]">
      <CardHeader>
        <CardTitle>Repository watcher</CardTitle>
        <CardDescription>
          You have {state.openedRepos.length} opened repositories.
        </CardDescription>
      </CardHeader>
      <CardContent class="grid gap-4">
        <div>
          <For each={state.openedRepos}>{(repo) => <RepoItem {...repo} />}</For>
        </div>
      </CardContent>
      <CardFooter>
        <Button class="w-full" onClick={openRepository}>
          <FolderUp class="mr-2 size-4" /> Open Repository
        </Button>
      </CardFooter>
    </Card>
  );
}

const RepoItem = (repo: { Path: string; Name: string }) => {
  const [checked, setChecked] = createSignal(false);
  const [_, { toggleWatching, removeRepository }] = useStore();

  createEffect(() => {
    toggleWatching(repo.Path, checked());
  });

  const handleRemove = () => {
    // First ensure watching is stopped
    if (checked()) {
      setChecked(false);
    }
    removeRepository(repo.Path);
  };

  return (
    <div class="mb-4 grid grid-cols-4 items-start pb-4 last:mb-0 last:pb-0">
      <div class="space-y-1 col-span-3">
        <p class="text-sm font-medium leading-none">{repo.Name}</p>
        <p class="text-sm text-muted-foreground">{repo.Path}</p>
      </div>
      <div class="col-span-1 flex items-center justify-end space-x-3 h-full">
        <Switch checked={checked()} onChange={setChecked}>
          <SwitchControl>
            <SwitchThumb />
          </SwitchControl>
        </Switch>
        <X
          class="w-5 h-5 mb-1 text-red-500 cursor-pointer hover:text-red-600"
          onClick={handleRemove}
        />
      </div>
    </div>
  );
};
