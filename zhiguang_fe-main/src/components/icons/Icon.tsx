import type { SVGProps } from "react";

export type IconProps = SVGProps<SVGSVGElement>;

const createIcon = (path: JSX.Element) => (props: IconProps) => (
  <svg
    width={props.width ?? 24}
    height={props.height ?? 24}
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    strokeWidth={props.strokeWidth ?? 1.6}
    strokeLinecap="round"
    strokeLinejoin="round"
    {...props}
  >
    {path}
  </svg>
);

export const HomeIcon = createIcon(
  <>
    <path d="M3.6 10.3 11.1 3.8a1.3 1.3 0 0 1 1.7 0l7.6 6.5" />
    <path d="M5.5 9.5V19a1 1 0 0 0 1 1h10.9a1 1 0 0 0 1-1v-9.4" />
    <path d="M9.5 20V13h5v7" />
  </>
);

export const SearchIcon = createIcon(
  <>
    <circle cx="10.6" cy="10.6" r="5.8" />
    <path d="m16.2 16.2 4.2 4.2" />
  </>
);

export const CreateIcon = createIcon(
  <>
    <path d="M5.5 19v-5.1a2 2 0 0 1 2-2h3.3" />
    <path d="M5.5 19h4.6" />
    <path d="M13.6 5.4c.8-.8 2-.8 2.8 0l1.7 1.7c.8.8.8 2 0 2.8L12 15l-3.4.9.9-3.4z" />
  </>
);

export const StudyIcon = createIcon(
  <>
    <path d="M4 7.5 12 4l8 3.5-8 3.6z" />
    <path d="m6.8 16.4 5.2-2.2 5.2 2.2M20 7.6V16" />
    <path d="m4 7.6 0 8.4a2 2 0 0 0 1.2 1.8l5.6 2.4" />
  </>
);

export const ProfileIcon = createIcon(
  <>
    <circle cx="12" cy="8.3" r="3.2" />
    <path d="M6.2 19.6a6.6 6.6 0 0 1 11.6 0" />
  </>
);

export const SparkIcon = createIcon(
  <>
    <path d="m12 2.8 1.5 4.6h4.9l-3.9 2.9 1.4 4.7L12 12.7 8.1 15l1.4-4.7-3.9-2.9h4.9z" />
  </>
);

export const ArrowRightIcon = createIcon(
  <>
    <path d="m10 7 5 5-5 5" />
    <path d="M5 12h10" />
  </>
);

export const HeartIcon = createIcon(
  <>
    <path d="M12 19c-4.8-2.7-7.5-5.6-7.5-9a4.5 4.5 0 0 1 8.2-2.4h.6A4.5 4.5 0 0 1 19.5 10c0 3.4-2.7 6.3-7.5 9z" />
  </>
);

export const BookmarkIcon = createIcon(
  <>
    <path d="M7.5 4h9a1 1 0 0 1 1 1v14l-5.5-3.5L6.5 19V5a1 1 0 0 1 1-1z" />
  </>
);

export const EyeIcon = createIcon(
  <>
    <path d="M2 12s3.6-7 10-7 10 7 10 7-3.6 7-10 7-10-7-10-7Z" />
    <circle cx="12" cy="12" r="3" />
  </>
);

export const MessageCircleIcon = createIcon(
  <>
    <path d="M21 11.5a8.38 8.38 0 0 1-.9 3.8 8.5 8.5 0 0 1-7.6 4.7 8.38 8.38 0 0 1-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 0 1-.9-3.8 8.5 8.5 0 0 1 4.7-7.6 8.38 8.38 0 0 1 3.8-.9h.5a8.48 8.48 0 0 1 8 8v.5z" />
  </>
);

export const ShareIcon = createIcon(
  <>
    <circle cx="18" cy="5" r="3" />
    <circle cx="6" cy="12" r="3" />
    <circle cx="18" cy="19" r="3" />
    <line x1="8.59" y1="13.51" x2="15.42" y2="17.49" />
    <line x1="15.41" y1="6.51" x2="8.59" y2="10.49" />
  </>
);

export const BotIcon = createIcon(
  <>
    <rect x="3" y="11" width="18" height="10" rx="2" />
    <circle cx="12" cy="5" r="2" />
    <path d="M12 7v4" />
    <line x1="8" y1="16" x2="8" y2="16" />
    <line x1="16" y1="16" x2="16" y2="16" />
  </>
);

export const CheckIcon = createIcon(
  <>
    <polyline points="20 6 9 17 4 12" />
  </>
);

export const CopyIcon = createIcon(
  <>
    <rect x="9" y="9" width="13" height="13" rx="2" ry="2" />
    <path d="M5 15H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9a2 2 0 0 1 2 2v1" />
  </>
);
