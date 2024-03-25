import { create } from "zustand";

interface Auth {
  isAuthenticated: boolean;
  setIsAuthenticated: (state: boolean) => void;
}

const useAuth = create<Auth>((set) => ({
  isAuthenticated: true,
  setIsAuthenticated: (state: boolean) => set({ isAuthenticated: state }),
}));

export default useAuth;
