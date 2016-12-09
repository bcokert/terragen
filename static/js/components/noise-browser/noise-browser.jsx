"use strict";

var React = require("react");
var Plotly = require("plotly.js");
var Ajax = require("../../ajax/ajax");
var TextField = require("../control/text-field/text-field.jsx");

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

        this.getLayout = this.getLayout.bind(this);
        this.getPlot = this.getPlot.bind(this);
        this.getConfig = this.getConfig.bind(this);

        this.fetchNoise = this.fetchNoise.bind(this);

        this.onChangeFrom = this.onChangeFrom.bind(this);
        this.onChangeTo = this.onChangeTo.bind(this);
        this.onChangeResolution = this.onChangeResolution.bind(this);

        this.componentDidMount = this.componentDidMount.bind(this);
        this.componentDidUpdate = this.componentDidUpdate.bind(this);
    }

    componentDidMount() {
        Plotly.plot(this.plotArea, [this.getPlot()], this.getLayout(), this.getConfig());
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

        if (
            previousState.t1 !== this.state.t1
            || previousState.t2 !== this.state.t2
            || previousState.t3 !== this.state.t3
            || previousState.value !== this.state.value
        ) {
            Plotly.deleteTraces(this.plotArea, 0);
            Plotly.addTraces(this.plotArea, this.getPlot());
        }
    }

    getLayout() {
        switch(this.props.dimension) {
            case 1:
                return {
                    margin: {l: 0, r: 0, t: 0, b: 0},
                    height: 300,
                    width: this.plotArea.offsetWidth
                };
            case 2:
                return {
                    margin: {l: 0, r: 0, t: 0, b: 0},
                    height: 700,
                    width: this.plotArea.offsetWidth
                };
            default:
                throw new Error("Unsupported dimension for noise browser, in getLayout");
        }
    }

    getConfig() {
        switch(this.props.dimension) {
            case 1:
                return {
                    showLink: false,
                    displayModeBar: false
                };
            case 2:
                return {
                    showLink: false,
                    displayModeBar: false,
                    scrollZoom: false
                };
            default:
                throw new Error("Unsupported dimension for noise browser, in getConfig");
        }
    }

    getPlot() {
        switch(this.props.dimension) {
            case 1:
                return {
                    x: this.state.t1,
                    y: this.state.value,
                    mode: "lines"
                };
            case 2:
                return {
                    x: this.state.t1,
                    y: this.state.t2,
                    z: this.state.value,
                    type: "mesh3d"
                };
            default:
                throw new Error("Unsupported dimension for noise browser, in getPlot");
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
        return this.state.errors.map(e => <span className="-error x-value">{e}</span>);
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
                <div className="-plotArea" ref={node => this.plotArea = node}></div>
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
