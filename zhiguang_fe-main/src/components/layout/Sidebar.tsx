import { NavLink } from "react-router-dom";
import { CreateIcon, HomeIcon, ProfileIcon, SearchIcon, StudyIcon } from "@/components/icons/Icon";
import styles from "./Sidebar.module.css";

const navItems = [
  { to: "/", label: "发现", Icon: HomeIcon },
  { to: "/search", label: "搜索", Icon: SearchIcon },
  { to: "/create", label: "创作", Icon: CreateIcon },
  { to: "/learn", label: "学习", Icon: StudyIcon },
  { to: "/profile", label: "我的", Icon: ProfileIcon }
] as const;

const Sidebar = () => {
  return (
    <aside className={styles.sidebar}>
      <div className={styles.brandHeader}>
        <span className={styles.brandTitle}>知光</span>
        <span className={styles.brandSubtitle}>ZHIGUANG</span>
      </div>

      <nav className={styles.nav}>
        {navItems.map(({ to, label, Icon }) => (
          <NavLink
            key={to}
            to={to}
            end={to === "/"}
            className={({ isActive }) => (isActive ? `${styles.link} ${styles.linkActive}` : styles.link)}
          >
            <Icon />
            <span className={styles.linkLabel}>{label}</span>
          </NavLink>
        ))}
      </nav>

      <div className={styles.footer}>
        <div className={styles.footerTag}>
          <span className={styles.statusDot} />
          <span>让知识发光</span>
        </div>
      </div>
    </aside>
  );
};

export default Sidebar;
