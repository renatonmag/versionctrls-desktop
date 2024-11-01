import { SelectFolder, StartWatcher } from "../../wailsjs/go/backend/App";

export const createUtils = (state, actions, setState) => {
  Object.assign(actions, {
    selectFolder: () => {
      SelectFolder().then((path) => {
        setState("folderPath", path);
      });
    },
    startWatching: () => {
      console.log("startWatching");
      StartWatcher(state.folderPath);
    },
  });
};
