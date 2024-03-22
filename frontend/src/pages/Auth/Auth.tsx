import AuthForm from "../../components/Header/Auth/AuthForm";
import styles from "./Auth.module.scss";
type Props = {};

const Auth = (props: Props) => {
  return (
    <div className={styles.container}>
      <AuthForm />
    </div>
  );
};

export default Auth;
