import { BrowserRouter, Route, Routes } from "react-router-dom";
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
        <Routes>
          <Route path="/" element={isAuthenticated ? <Main /> : <Auth />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
