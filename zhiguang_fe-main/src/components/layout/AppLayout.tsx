import type { PropsWithChildren, ReactNode } from "react";
import Sidebar from "./Sidebar";
import styles from "./AppLayout.module.css";

type AppLayoutProps = PropsWithChildren<{
  header?: ReactNode;
  variant?: "default" | "detail" | "plain" | "cardless";
}>;

const AppLayout = ({ children, header, variant = "default" }: AppLayoutProps) => {
  return (
    <div className="app-shell">
      <Sidebar />
      <div className={`${styles.mainContent} ${variant === "detail" || variant === "cardless" ? styles.mainDetail : ""}`}>
        {header ? <div className={styles.headerWrapper}>{header}</div> : null}
        <main className={styles.body}>{children}</main>
      </div>
    </div>
  );
};

export default AppLayout;
