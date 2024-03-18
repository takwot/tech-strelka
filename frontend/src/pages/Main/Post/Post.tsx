import styles from "./Post.module.scss";

type Props = {};

const Post = (props: Props) => {
  return (
    <div className={styles.container}>
      <div className={styles.top_container}>
        <div className={styles.avatar}></div>
        <p>NICKNAME</p>
      </div>
      <div className={styles.photo_cont}></div>
      <div className={styles.tags}>
        <p>#bus</p>
        <p>#bus</p>
        <p>#bus</p>
        <p>#bus</p>
        <p>#bus</p>
        <p>#bus</p>
        <p>#bus</p>
        <p>#bus</p>
        <p>#bus</p>
        <p>#bus</p>
        <p>#bus</p>
        <p>#bus</p>
      </div>
      <div className={styles.data}>
        <p>21.09.2024</p>
      </div>
    </div>
  );
};

export default Post;
