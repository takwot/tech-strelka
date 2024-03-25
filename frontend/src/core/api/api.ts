import axios from "axios";

const instanse = axios.create({
  baseURL: "http://localhost:5000/api",
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
  createPhoto(form: FormData, token: string) {
    return instanse.post("/photo?tags=[1, 2, 3]", form, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  },
};

export default Api;
