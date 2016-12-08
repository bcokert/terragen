"use strict";

var React = require("react");
var classNames = require("classnames");

require("./button.less");

class Button extends React.Component {
    constructor(props) {
        super(props);
    }

    render() {
        const classes = classNames("Button", this.props.className, {});

        return (
            <button className={classes} onClick={this.props.onClick}>{this.props.children}</button>
        );
    }
}

Button.displayName = "Button";

Button.propTypes = {
    children: React.PropTypes.node.isRequired,
    className: React.PropTypes.string,
    onClick: React.PropTypes.func.isRequired
};

Button.defaultProps = {
    className: ""
};

module.exports = Button;
