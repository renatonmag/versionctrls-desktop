import { createContext, useContext, ParentComponent } from "solid-js";
import { createStore } from "solid-js/store";
import { createUtils } from "./createUtils";

// Define the store state type
type StoreState = {
  // Add your state properties here
  count: number;
  // ... other state properties
};

// Create the context
const StoreContext = createContext<[StoreState, any]>();

// Create the provider component
export const StoreProvider: ParentComponent = (props) => {
  const [state, setState] = createStore<StoreState>({
    count: 0,
    folderPath: "",
  });
  const actions = {};
  const store = [state, actions];
  createUtils(state, actions, setState);

  return (
    <StoreContext.Provider value={store}>
      {props.children}
    </StoreContext.Provider>
  );
};

// Create a hook to use the store
export const useStore = () => {
  const context = useContext(StoreContext);
  if (!context) {
    throw new Error("useStore must be used within a StoreProvider");
  }
  return context;
};
