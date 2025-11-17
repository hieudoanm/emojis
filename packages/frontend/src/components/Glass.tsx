import {
  FC,
  ReactNode,
  ButtonHTMLAttributes,
  InputHTMLAttributes,
  TextareaHTMLAttributes,
  SelectHTMLAttributes,
} from 'react';

const baseSharedClasses = `
  border border-white/20
  bg-white/10
  backdrop-blur-lg
  shadow-lg shadow-black/30
  text-white placeholder:text-neutral-400
  transition duration-300
  focus:outline-none focus:ring-2 focus:ring-white/30
`;

/* ---------------- BUTTON ---------------- */
const Button: FC<
  {
    children: ReactNode;
    className?: string;
  } & ButtonHTMLAttributes<HTMLButtonElement>
> = ({ type = 'button', children, className = '', disabled, ...props }) => {
  return (
    <button
      type={type}
      disabled={disabled}
      className={`inline-flex items-center justify-center rounded-full px-6 py-2 font-semibold ${baseSharedClasses} ${
        disabled ? 'cursor-not-allowed opacity-50' : 'cursor-pointer'
      } ${className}`}
      {...props}>
      {children}
    </button>
  );
};

/* ---------------- CARD ---------------- */
const Card: FC<{ children: ReactNode; className?: string }> = ({
  children,
  className = '',
}) => {
  return (
    <div
      className={`relative overflow-hidden rounded-2xl border border-white/10 bg-white/5 p-6 shadow-xl shadow-black/30 backdrop-blur-lg transition-all duration-300 sm:p-8 ${className}`}>
      {children}
    </div>
  );
};

/* ---------------- INPUT ---------------- */
const Input: FC<InputHTMLAttributes<HTMLInputElement>> = ({
  className = '',
  ...props
}) => {
  return (
    <input
      className={`rounded-full px-4 py-2 ${baseSharedClasses} ${className}`}
      {...props}
    />
  );
};

/* ---------------- TEXTAREA ---------------- */
const TextArea: FC<TextareaHTMLAttributes<HTMLTextAreaElement>> = ({
  className = '',
  ...props
}) => {
  return (
    <textarea
      rows={4}
      className={`resize-none rounded-xl px-4 py-3 ${baseSharedClasses} ${className}`}
      {...props}
    />
  );
};

/* ---------------- SELECT ---------------- */
const Select: FC<SelectHTMLAttributes<HTMLSelectElement>> = ({
  className = '',
  children,
  ...props
}) => {
  return (
    <select
      className={`rounded-full bg-white/10 px-4 py-2 text-white ${baseSharedClasses} appearance-none ${className}`}
      {...props}>
      {children}
    </select>
  );
};

/* ---------------- BADGE ---------------- */
const Badge: FC<{ children: ReactNode; className?: string }> = ({
  children,
  className = '',
}) => (
  <span
    className={`inline-block rounded-full border border-white/20 bg-white/10 px-3 py-1 text-sm font-medium shadow shadow-black/20 backdrop-blur-lg ${className}`}>
    {children}
  </span>
);

/* ---------------- MODAL ---------------- */
const Modal: FC<{
  open: boolean;
  onClose: () => void;
  children: ReactNode;
}> = ({ open, onClose, children }) =>
  !open ? null : (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black/40 backdrop-blur-sm">
      <div className="relative rounded-2xl border border-white/20 bg-white/10 p-6 shadow-lg backdrop-blur-xl">
        {children}
        <button
          onClick={onClose}
          className="absolute top-3 right-3 text-white hover:text-neutral-300"
          aria-label="Close">
          ✕
        </button>
      </div>
    </div>
  );

/* ---------------- TOOLTIP ---------------- */
const Tooltip: FC<{ label: string; children: ReactNode }> = ({
  label,
  children,
}) => (
  <div className="group relative">
    {children}
    <span className="absolute bottom-full mb-2 hidden rounded-md bg-black/70 px-2 py-1 text-xs text-white group-hover:block">
      {label}
    </span>
  </div>
);

/* ---------------- CHECKBOX ---------------- */
const Checkbox: FC<InputHTMLAttributes<HTMLInputElement>> = ({
  className = '',
  ...props
}) => (
  <input
    type="checkbox"
    className={`h-4 w-4 rounded border-white/30 bg-white/10 checked:bg-white/30 focus:ring-white/30 ${className}`}
    {...props}
  />
);

/* ---------------- RADIO ---------------- */
const Radio: FC<InputHTMLAttributes<HTMLInputElement>> = ({
  className = '',
  ...props
}) => (
  <input
    type="radio"
    className={`h-4 w-4 rounded-full border-white/30 bg-white/10 checked:bg-white/30 focus:ring-white/30 ${className}`}
    {...props}
  />
);

/* ---------------- SWITCH ---------------- */
const Switch: FC<{
  checked: boolean;
  onChange: () => void;
  className?: string;
}> = ({ checked, onChange, className = '' }) => (
  <div
    role="button"
    className={`relative h-6 w-12 cursor-pointer rounded-full ${checked ? 'bg-white/30' : 'bg-white/10'} ${className}`}
    onClick={onChange}>
    <div
      className={`absolute top-1 left-1 h-4 w-4 rounded-full bg-white shadow-md transition-transform ${
        checked ? 'translate-x-6' : ''
      }`}
    />
  </div>
);

/* ---------------- AVATAR ---------------- */
const Avatar: FC<{
  src: string;
  alt?: string;
  size?: number;
  className?: string;
}> = ({ src, alt, size = 40, className = '' }) => (
  <img
    src={src}
    alt={alt}
    width={size}
    height={size}
    className={`rounded-full border border-white/20 shadow shadow-black/20 ${className}`}
  />
);

export const Glass = {
  Button,
  Card,
  Input,
  TextArea,
  Select,
  Badge,
  Modal,
  Tooltip,
  Checkbox,
  Radio,
  Switch,
  Avatar,
};
