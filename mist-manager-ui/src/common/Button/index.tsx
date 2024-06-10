import React from "react";

interface ButtonProps {
  className?: string;
  onClick: () => void;
  children: React.ReactNode;
}

const Button = ({ className, onClick, children }: ButtonProps): JSX.Element => {
  return (
    <button className={className} onClick={onClick}>
      {children}
    </button>
  );
};

export default Button;
