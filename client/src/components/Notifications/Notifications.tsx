import React from "react";

interface INotificationsProps<T> {
  timeout: number;
  children: (notifications: T[], notify: (x: T) => void) => React.ReactNode;
}

interface INotificationsState<T> {
  notifications: { notification: T; created: number }[];
}

class Notifications<T> extends React.Component<
  INotificationsProps<T>,
  INotificationsState<T>
> {
  vacuumInterval = 0;

  state: INotificationsState<T> = {
    notifications: []
  };

  componentDidMount(): void {
    this.vacuumInterval = window.setInterval(this.vacuumNotifications, 200);
  }

  componentWillUnmount(): void {
    this.vacuumInterval && clearInterval(this.vacuumInterval);
  }

  vacuumNotifications = () => {
    const { notifications } = this.state;
    const i = notifications.findIndex(
      ({ created }) => created < Date.now() - this.props.timeout
    );
    if (i != -1)
      this.setState({
        notifications: notifications.slice(0, i) // drop notifications older then now - timeout
      });
  };

  notify = (notification: T) => {
    this.setState(({ notifications }) => ({
      notifications: [
        {
          notification,
          created: Date.now()
        },
        ...notifications
      ]
    }));
  }

  render() {
    return this.props.children(
      this.state.notifications.map(x => x.notification),
      this.notify
    );
  }
}

export default Notifications;
