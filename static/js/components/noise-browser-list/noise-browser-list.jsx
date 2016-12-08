"use strict";

var React = require("react");
var NoiseBrowser = require("../noise-browser/noise-browser.jsx");
var TextField = require("../control/text-field/text-field.jsx");

require("./noise-browser-list.less");

class NoiseBrowserList extends React.Component {
    constructor(props) {
        super(props);
    }

    renderNoiseBrowsers() {
        return this.props.browserList.map((browserData, i) => <NoiseBrowser dimension={browserData.dimension} displayName={browserData.displayName} endpoint={browserData.endpoint} initialNoiseFunction={browserData.noiseFunction} key={"NoiseBrowser1D_" + i}/>);
    }

    render() {
        return (
            <div className="NoiseBrowserList">
                <div className="-header">
                    <span className="-title">{this.props.presetCollectionName}</span>
                    <span className="-description">{this.props.description}</span>
                    <div className="-componentList">
                        <TextField label="NoiseFunction" readOnly value={this.props.noiseFunction}/>
                        <TextField label="Generator" readOnly value={this.props.generator}/>
                        <TextField label="Transformer" readOnly value={this.props.transformer}/>
                        <TextField label="Synthesizer" readOnly value={this.props.synthesizer}/>
                    </div>
                </div>
                {this.renderNoiseBrowsers()}
            </div>
        );
    }
}

NoiseBrowserList.displayName = "NoiseBrowserList";

NoiseBrowserList.propTypes = {
    browserList: React.PropTypes.arrayOf(React.PropTypes.shape({
        dimension: React.PropTypes.number.isRequired,
        displayName: React.PropTypes.string.isRequired,
        endpoint: React.PropTypes.string.isRequired,
        noiseFunction: React.PropTypes.string.isRequired
    })).isRequired,
    description: React.PropTypes.string.isRequired,
    generator: React.PropTypes.string,
    noiseFunction: React.PropTypes.string,
    presetCollectionName: React.PropTypes.string.isRequired,
    synthesizer: React.PropTypes.string,
    transformer: React.PropTypes.string
};

NoiseBrowserList.defaultProps = {
    generator: "None",
    noiseFunction: "N/A",
    synthesizer: "None",
    transformer: "None"
};

module.exports = NoiseBrowserList;
