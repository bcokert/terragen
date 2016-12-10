"use strict";

var React = require("react");
var Ajax = require("../../ajax/ajax");
var TextField = require("../control/text-field/text-field.jsx");
var LinePlotArea = require("../line-plot-area/line-plot-area");

require("./noise-browser.less");

class NoiseBrowser extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            t1: [],
            t2: [],
            t3: [],
            value: [],

            from: [0,0,0].slice(0,props.dimension).join(","),
            to: [5,2,2].slice(0,props.dimension).join(","),
            resolution: String([40, 20][props.dimension-1]),
            noiseFunction: props.initialNoiseFunction,

            errors: []
        };

        this.fetchNoise = this.fetchNoise.bind(this);

        this.onChangeFrom = this.onChangeFrom.bind(this);
        this.onChangeTo = this.onChangeTo.bind(this);
        this.onChangeResolution = this.onChangeResolution.bind(this);

        this.componentDidMount = this.componentDidMount.bind(this);
        this.componentDidUpdate = this.componentDidUpdate.bind(this);
    }

    componentDidMount() {
        this.fetchNoise();
    }

    componentDidUpdate(previousProps, previousState) {
        if (
            previousState.from !== this.state.from
            || previousState.to !== this.state.to
            || previousState.resolution !== this.state.resolution
            || previousState.noiseFunction !== this.state.noiseFunction
        ) {
            this.fetchNoise();
        }
    }

    fetchNoise() {
        Ajax.request({
            url: this.props.endpoint,
            method: "GET",
            queryParams: {
                from: this.state.from,
                to: this.state.to,
                resolution: this.state.resolution,
                noiseFunction: this.state.noiseFunction
            }
        }).then(response => {
            if (response.rawNoise && typeof response.rawNoise === "object") {
                this.setState({
                    t1: response.rawNoise.t1 || [],
                    t2: response.rawNoise.t2 || [],
                    t3: response.rawNoise.t3 || [],
                    value: response.rawNoise.value || [],
                    errors: []
                });
            } else {
                throw new Error("Invalid response from server: " + JSON.stringify(response));
            }
        }).catch(e => {
            this.setState({ 
                errors: [e.message]
            });
        });
    }

    renderErrors() {
        return this.state.errors.map(e => <span className="-error x-value" key={e} >{e}</span>);
    }

    onChangeFrom(newFrom) {
        this.setState({from: newFrom});
    }

    onChangeTo(newTo) {
        this.setState({to: newTo});
    }

    onChangeResolution(newResolution) {
        this.setState({resolution: newResolution});
    }

    render() {
        return (
            <div className="NoiseBrowser">
                <div className="-title">
                    <span className="x-label">{this.props.displayName}</span>
                </div>
                <div className="-control -top">
                    <TextField label="From" onChange={this.onChangeFrom} validate={v => v.split(",").length === this.props.dimension && v.split(",").map(n => Math.floor(parseFloat(n))).join(",").length === v.length} value={this.state.from}/>
                    <TextField label="To" onChange={this.onChangeTo} validate={v => v.split(",").length === this.props.dimension && v.split(",").map(n => Math.floor(parseFloat(n))).join(",").length === v.length} value={this.state.to}/>
                    <TextField label="Resolution" onChange={this.onChangeResolution} validate={v => !isNaN(v) && String(parseInt(v)).length === v.length} value={this.state.resolution}/>
                    <TextField label="NoiseFunction" readOnly value={this.state.noiseFunction}/>
                </div>
                <div className="-errors">
                    {this.renderErrors()}
                </div>
                <LinePlotArea height={300} width={window.innerWidth - 40} x={this.state.t1} y={this.state.value} />
                <div className="-control -bottom"></div>
            </div>
        );
    }
}

NoiseBrowser.displayName = "NoiseBrowser";

NoiseBrowser.propTypes = {
    dimension: React.PropTypes.number.isRequired,
    displayName: React.PropTypes.string.isRequired,
    endpoint: React.PropTypes.string.isRequired,
    initialNoiseFunction: React.PropTypes.string.isRequired
};

module.exports = NoiseBrowser;
