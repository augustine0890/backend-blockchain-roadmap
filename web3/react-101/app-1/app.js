import React from "react";

export default class App extends React.Component {
  render() {
    const users = [
      { name: 'Robin'},
      { name: 'Markus'},
    ];

    return (
      <ul>
        {users.map(user => <li>{user.name}</li>)}
      </ul>
    )
  }
}