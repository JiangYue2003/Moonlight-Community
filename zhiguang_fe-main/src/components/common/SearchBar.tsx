import { useState, useRef, type ChangeEvent } from "react";
import { SearchIcon } from "@/components/icons/Icon";
import styles from "./SearchBar.module.css";

type SearchBarProps = {
  placeholder?: string;
  value: string;
  onChange: (value: string) => void;
  onSubmit?: () => void;
  buttonLabel?: string;
  suggestions?: string[];
  suggestLoading?: boolean;
  onSuggestionClick?: (value: string) => void;
};

const SearchBar = ({
  placeholder = "搜索知文、标签、作者...",
  value,
  onChange,
  onSubmit,
  buttonLabel = "搜索",
  suggestions = [],
  suggestLoading = false,
  onSuggestionClick
}: SearchBarProps) => {
  const [focused, setFocused] = useState(false);
  const inputRef = useRef<HTMLInputElement>(null);

  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    onChange(event.target.value);
  };

  return (
    <div className={styles.wrapper}>
      <input
        ref={inputRef}
        className={styles.input}
        value={value}
        placeholder={placeholder}
        onChange={handleChange}
        onFocus={() => setFocused(true)}
        onBlur={() => setTimeout(() => setFocused(false), 140)}
        onKeyDown={(e) => {
          if (e.key === "Enter") {
            onSubmit?.();
          }
        }}
      />
      {!focused && !value && (
        <span className={styles.kbd}>⌘K</span>
      )}
      <button
        className={styles.button}
        type="button"
        onClick={onSubmit}
        aria-label={buttonLabel}
        title={buttonLabel}
      >
        <SearchIcon width={18} height={18} strokeWidth={2} />
      </button>

      {focused && (value?.trim()?.length ?? 0) > 0 && (
        <div className={styles.dropdown}>
          {suggestLoading ? (
            <div className={styles.dropdownEmpty}>正在检索相关建议...</div>
          ) : suggestions?.length ? (
            suggestions.map((s) => (
              <div
                key={s}
                className={styles.dropdownItem}
                onMouseDown={() => onSuggestionClick?.(s)}
              >
                <SearchIcon width={14} height={14} strokeWidth={1.8} style={{ opacity: 0.5 }} />
                <span>{s}</span>
              </div>
            ))
          ) : (
            <div className={styles.dropdownEmpty}>无联想结果</div>
          )}
        </div>
      )}
    </div>
  );
};

export default SearchBar;
