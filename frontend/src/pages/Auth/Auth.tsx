import AuthForm from "../../components/Header/Auth/AuthForm";
import styles from "./Auth.module.scss";

const Auth = () => {
  return (
    <div className={styles.container}>
      <AuthForm />
    </div>
  );
};

export default Auth;
