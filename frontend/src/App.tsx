import { createSignal } from "solid-js";
import { Button } from "./components/ui/button";
import { useStore } from "./store/store";

export default function App() {
  const [_, { selectFolder, startWatching }] = useStore();

  return (
    <div class="flex justify-center h-screen">
      <Button onClick={selectFolder}>Select Folder</Button>
      <Button onClick={startWatching}>Watch</Button>
    </div>
  );
}
