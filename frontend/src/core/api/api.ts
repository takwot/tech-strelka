import axios from "axios";

const instanse = axios.create({
  baseURL: "http://45.84.225.194:3000/api",
  // baseURL: "http://localhost:3000/api",
  withCredentials: false,
});

const Api = {
  login(nickname: string, password: string) {
    return instanse.post("/auth/login", {
      name: nickname,
      password,
    });
  },
  register(username: string, password: string, email: string) {
    return instanse.post("/auth/register", {
      name: username,
      email,
      password,
    });
  },
  createPhoto(form: FormData, id: number) {
    return instanse.post(`/photo?id=${id}`, form, {
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
    });
  },
  getAllPhoto(id: number) {
    return instanse.get(`/photo/all?id=${id}`);
  },
  deletePhoto(filename: string) {
    return instanse.delete(`/photo?filename=${filename}`);
  },
  createAlbum(name: string, author: number, photos: string[]) {
    return instanse.post("/album", {
      name,
      author,
      photos,
    });
  },
  deleteAlbum(id: number) {
    return instanse.delete(`/album?id=${id}`);
  },
  getAllAlbum(id: number) {
    return instanse.get(`/album?author=${id}`);
  },
};

export default Api;
