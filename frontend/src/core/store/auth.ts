import { create } from "zustand";

type User = {
  email: string;
  id: number;
  name: string;
};

interface Auth {
  isAuthenticated: boolean;
  setIsAuthenticated: (state: boolean) => void;
  user: User;
  setUser: (user: User) => void;
}

const useAuth = create<Auth>((set) => ({
  isAuthenticated: false,
  setIsAuthenticated: (state: boolean) => set({ isAuthenticated: state }),
  user: { email: "", id: 0, name: "" },
  setUser: (user: User) => set({ user }),
}));

export default useAuth;
