# Nearby
a social network app based on geo location

Let's learn [React](https://github.com/facebookincubator/create-react-app)! 

## What Is React And How Do We Use It?  
`React` is a JavaScript library for building `user interfaces` (the page you are currently looking at and interacting with) .

The benefits of using React include:
* `Components-based`: Build encapsulated components that manage their own state, then compose them to make complex UIs 
* `Declarative`: React makes it painless to create interactive UIs.
* `Learn Once, Write Anywhere`: You can develop new features in React without rewriting existing code.

## Table of Contents

- [Component and Props](#component-and-props)


## Component and Props

In the traditional HTML and Javascript world, web pages are divided and represented in blocks, by tags like `div` or `section`. Then we would have some JS functions to control the behavior of these tags. 

In React, the developer took up a notch, they kind combined these two concepts together, making writing dynamic web page seemingless and easier.  

`Component` let you split the UI into **independent**, **reusable** pieces. Conceptually, components are like JavaScript functions. 
They accept arbitrary inputs (called “props”) and return React elements describing what should appear on the screen.

#### Functional and Class Components
The simplest way to define a component is to write a Javascript function:
```
function Welcome(props) { 
     return <h1>Hello, {props.name}</h1>; 
}
``` 
This function is a valid React component because it accpets a single `props` (which stands for properties) object argument with data and returns a React component. We call such components "functional" because, they are functions.

Or to define a `class component`:
```
class Welcome extends React.Component {
    render() {
        return <h1> Hello, {this.props.name}</h1>;
    }
}
```
#### Rendering a Component
When React sees an element representing a user-defined component, it passes JSX attributes to this component as a single object. We call this object “props”. For example, this code renders “Hello, Sara” on the page:
```
function Welcome(props) {
  return <h1>Hello, {props.name}</h1>;
}

const element = <Welcome name="Sara" />;
ReactDOM.render(
  element,
  document.getElementById('root')
);
```
Wait a minute, what is this? 
Let's recap what actually happens in this example:
* We call `ReactDOM.render()` with the `<Welcome name="Sara" />` element. 
* React calls the Welcome component with {name: 'Sara'} as the props
* Our Welcome component returns a `<h1>Hello, Sara</h1>` element as the result.
* React DOM efficiently updates the DOM to match `<h1>Hello, Sara</h1>`.

#### Composing Components
Components can refer to other components in their output. This lets us use the same component abstraction for any level of detail. A button, a form, a dialog, a screen: in React apps, all those are commonly expressed as components. Let's use the `Welcome` example again:
```
function Welcome(props) {
 return <h1>Hello, {props.name}</h1>;
}

function App() {
 return (
   <div>
     <Welcome name="Sara" />
     <Welcome name="John" />
     <Welcome name="Richard" />
   </div>
 );
}

ReactDOM.render(
 <App />,
 document.getElementById('root')
);
```
#### Extracting Components 
Now we have some sense of how to congregate components together, we can also break components into smaller ones.
For example, consider this large component, can we extract it? 
```
function Comment(props) {
  return (
    <div className="Comment">
      <div className="UserInfo">
        <img className="Avatar"
          src={props.author.avatarUrl}
          alt={props.author.name}
        />
        <div className="UserInfo-name">
          {props.author.name}
        </div>
      </div>
      <div className="Comment-text">
        {props.text}
      </div>
      <div className="Comment-date">
        {formatDate(props.date)}
      </div>
    </div>
  );
}

```
**[Try it here first](https://codepen.io/pen?&editors=0010)**. Hint, it accepts author (an object), text (a string), and date (a date) as props, and describes a comment on a social media website. 

See the answer [here](https://codepen.io/pen?&editors=0010).








