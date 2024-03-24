import { create } from "zustand";

interface Auth {
  isAuthenticated: boolean;
  setIsAuthenticated: (state: boolean) => void;
}

const useAuth = create<Auth>((set) => ({
  isAuthenticated: false,
  setIsAuthenticated: (state: boolean) => set({ isAuthenticated: state }),
}));

export default useAuth;
