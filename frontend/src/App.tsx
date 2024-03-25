import { BrowserRouter, Navigate, Route, Routes } from "react-router-dom";
import Header from "./components/Header/Header";
import Main from "./pages/Main/Main";
import useAuth from "./core/store/auth";
import Auth from "./pages/Auth/Auth";

function App() {
  const { isAuthenticated } = useAuth();

  return (
    <>
      <BrowserRouter>
        <Header />
        {isAuthenticated ? <Navigate to={"/"} /> : <Navigate to={"/auth"} />}
        <Routes>
          <Route path="/auth" element={<Auth />} />
          <Route path="/" element={<Main />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
