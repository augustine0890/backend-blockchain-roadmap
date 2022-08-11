import React from "react";

const doFilter = query => user => query === user.name;

export default class App extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      query: '',
    };

    this.onChange = this.onChange.bind(this);
  }

  onChange(event) {
    this.setState({ query: event.target.value });
  }

  render() {
    const users = [
      { name: 'Robin'},
      { name: 'Markus'},
    ];

    return (
      <div>
        <ul>
          {
            users
              .filter(doFilter(this.state.query))
              .map(myuser => <li>{myuser.name}</li>)
          }
        </ul>
        <input
          type="text"
          onChange={this.onChange}
        />
      </div>
    )
  }
}