import { create } from "zustand";

interface AlbumState {
  create: boolean;
  setCreate: (state: boolean) => void;
}

const useAlbums = create<AlbumState>((set) => ({
  create: false,
  setCreate: (state: boolean) => set({ create: state }),
}));

export default useAlbums;
