import {
  ListRepos,
  OpenRepository,
  ToggleWatcher,
  VerifyIntegration,
  RemoveRepository,
} from "wailsjs/go/backend/App";

const getRepoName = (path: string) => {
  return path.split("/").pop() || "";
};

export const createGit = (state: any, actions: any, setState: any) => {
  Object.assign(actions, {
    openRepository: () => {
      OpenRepository().then((result) => {
        if (result.Error) {
          console.error(result.Error);
          return;
        }
        setState("openedRepos", [
          ...state.openedRepos,
          { Path: result.Path, Name: getRepoName(result.Path) },
        ]);
      });
    },
    toggleWatching: (path: string, watch: boolean) => {
      ToggleWatcher(path, watch);
    },
    verifyIntegration: async () => {
      const result = await VerifyIntegration(state.folderPath);
      console.log(result);
    },
    loadRepos: async () => {
      const result = await ListRepos();
      setState("openedRepos", result);
    },
    removeRepository: async (path: string) => {
      try {
        await RemoveRepository(path);
        // Remove from local state
        setState(
          "openedRepos",
          state.openedRepos.filter((repo: any) => repo.Path !== path)
        );
      } catch (error) {
        console.error("Error removing repository:", error);
      }
    },
  });
};
