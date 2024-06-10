import React from "react";

import "./_accordion.scss";
import { Button, Divider } from "../";

interface AccordionProps {
  header: JSX.Element;
  children: React.ReactNode;
  className?: string;
}

const Accordion = ({
  header,
  children,
  className,
}: AccordionProps): JSX.Element => {
  const [isOpen, setIsOpen] = React.useState(false);

  return (
    <div className={`${className} base-accordion-container`}>
      <div className="base-accordion-header">
        <Button onClick={() => setIsOpen(!isOpen)}>{isOpen ? "-" : "+"}</Button>
        {header}
      </div>
      {isOpen && <Divider />}
      {children}
    </div>
  );
};

export default Accordion;
