import React from "react";

import "./_card.scss";

interface CardProps {
  className?: string;
  children: React.ReactNode;
}

const Card = ({ className, children }: CardProps): JSX.Element => {
  return <div className={`common-card ${className}`}>{children}</div>;
};

export default Card;
