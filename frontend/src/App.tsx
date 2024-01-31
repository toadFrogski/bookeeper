import { FC, useContext } from "react";
import styles from "./app.module.scss";
import { SessionContext } from "./contexts/session";
import { Theme } from "./contexts/session/Context";

const App: FC = () => {
  const session = useContext(SessionContext);

  const handleTheme = () => {
    switch (session.theme) {
      case Theme.dark:
        session.setTheme(Theme.light);
        break;
      case Theme.light:
        session.setTheme(Theme.dark);
        break;
    }
  };

  return (
    <>
      <header className={styles.header} onClick={handleTheme}>
        <p>Test header</p>
      </header>
      <section className={`m-48 p-24 ${styles.body}`}>
        <h1>Test body</h1>
      </section>
      <footer className={styles.footer}>
        <p>Test footer</p>
      </footer>
    </>
  );
};

export default App;
