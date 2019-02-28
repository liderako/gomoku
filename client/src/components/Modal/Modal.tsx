import React from "react";
import enhanceWithClickOutside from "react-click-outside";
import styles from "./Modal.module.css";

interface IModalProps {
  children: React.ReactNode;
  onClose(): void;
}

const Box = enhanceWithClickOutside(
  class extends React.Component<IModalProps> {
    handleClickOutside = () => this.props.onClose();

    render() {
      const { children } = this.props;
      return <div className={styles.box}>{children}</div>;
    }
  }
);

const Modal: React.FunctionComponent<IModalProps> = ({ children, onClose }) => (
  <div className={styles.overlay}>
    <Box onClose={onClose}>{children}</Box>
  </div>
);

export default Modal;
