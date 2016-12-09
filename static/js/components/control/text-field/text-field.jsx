"use strict";

var React = require("react");
var _ = require("underscore");
var classNames = require("classnames");

require("./text-field.less");

class TextField extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            isFocused: false,
            isHovered: false,
            isValid: props.validate(props.value),
            value: props.value
        };

        this.onChange = this.onChange.bind(this);
    }

    componentWillReceiveProps(newProps) {
        if (newProps.value !== this.state.value) {
            this.setState({
                isValid: this.props.validate(newProps.value),
                value: newProps.value
            });
        }
    }

    componentDidMount() {
        this.onChangeDebounced = _.debounce(this.props.onChange, this.props.debounceTimeout);
    }

    onChange(e) {
        const newValue = e.target.value;
        this.setState({
            isValid: this.props.validate(newValue),
            value: newValue
        });
        this.onChangeDebounced(newValue);
    }

    render() {
        let additionalProps = {};
        if (this.props.readOnly) {
            additionalProps.readOnly = true;
        }

        const classes = classNames("TextField", this.props.className, {
            "-error": !this.state.isValid,
            "-focus": this.state.isFocused,
            "-readOnly": this.props.readOnly,
            "-hover": this.state.isHovered
        });
        return (
            <span className={classes}>
                <label className="-label" onMouseOut={() => this.setState({isHovered: false})} onMouseOver={() => this.setState({isHovered: true})}>
                    {this.props.label}
                </label>
                <input
                    {...additionalProps}
                    className={classNames("-input")}
                    onBlur={() => this.setState({isFocused: false})}
                    onChange={this.onChange}
                    onFocus={() => this.setState({isFocused: true})}
                    onMouseOut={() => this.setState({isHovered: false})}
                    onMouseOver={() => this.setState({isHovered: true})}
                    placeholder={this.props.placeholder}
                    type="text"
                    value={this.state.value}
                />
            </span>
        );
    }
}

TextField.displayName = "Text Field";

TextField.propTypes = {
    className: React.PropTypes.string,
    debounceTimeout: React.PropTypes.number,
    label: React.PropTypes.string.isRequired,
    onChange: React.PropTypes.func,
    placeholder: React.PropTypes.string,
    readOnly: React.PropTypes.bool,
    validate: React.PropTypes.func,
    value: React.PropTypes.string
};

TextField.defaultProps = {
    debounceTimeout: 150,
    placeholder: "",
    onChange: () => null,
    readOnly: false,
    validate: () => true,
    value: ""
};

module.exports = TextField;
