import React from "react";
import withClickOutside from "react-click-outside";
import DropdownStyles from "./Dropdown.module.css";

interface IDropdownProps {
  renderButton(isOpen: boolean, toggle: () => void): React.ReactNode;
  renderContent(): React.ReactNode;
}

interface IDropdownState {
  isOpen: boolean;
}

class Dropdown extends React.Component<IDropdownProps, IDropdownState> {
  state = {
    isOpen: false
  };

  close = () =>
    this.setState({
      isOpen: false
    });

  toggle = () => {
    this.setState(({ isOpen }) => ({
      isOpen: !isOpen
    }));
  };

  handleClickOutside = this.close;

  render() {
    const { renderButton, renderContent } = this.props;
    const { isOpen } = this.state;
    return (
      <div className={DropdownStyles.container}>
        {renderButton(isOpen, this.toggle)}
        {isOpen && (
          <div className={DropdownStyles.content}>{renderContent()}</div>
        )}
      </div>
    );
  }
}

export default withClickOutside(Dropdown);
