"use strict";

var React = require("react");

require("./line-plot-area.less");

class LinePlotArea extends React.Component {
    constructor(props) {
        super(props);

        this.canvas = null;

        this.componentDidMount = this.componentDidMount.bind(this);
        this.componentDidUpdate = this.componentDidUpdate.bind(this);

        this.updateCanvasNode = node => this.canvas = node;
    }

    componentDidMount() {
        this.componentDidUpdate();
    }

    componentDidUpdate() {
        // Get the context, and resize it in case the window has been resized
        const context = this.canvas.getContext("2d");
        if (context.canvas.width !== context.canvas.clientWidth || context.canvas.height !== context.canvas.clientHeight) {
            context.canvas.width = context.canvas.clientWidth;
            context.canvas.height = context.canvas.clientHeight;
        }
        context.clearRect(0, 0, context.canvas.width, context.canvas.height);

        // Short circuit if there's nothing to draw
        var values = this.props.y;
        if (values.length === 0) {
            return;
        }

        // Find the largest absolute value for scaling. We scale by the furthest value from the center
        var maxAbs = values[0];
        for (var i = 0; i < values.length; i++) {
            if (Math.abs(values[i]) > maxAbs) {
                maxAbs = Math.abs(values[i]);
            }
        }

        // Calculate graph sizes and scaling factors
        var logicalWidth = this.props.x[this.props.x.length - 1] - this.props.x[0];
        var xScaleFactor = this.props.width / logicalWidth;
        var halfHeight = this.props.height / 2;
        var yScaleFactor = halfHeight / (maxAbs * 1.03);

        // Draw a light line through the center
        context.lineWidth = 1;
        context.strokeStyle = "#ccc";
        context.beginPath();
        context.moveTo(0, halfHeight);
        context.lineTo(this.props.width, halfHeight);
        context.stroke();

        // Draw lines between the points pairwise, and a background line at the start of every interval
        for (i = 1; i < values.length; i++) {
            if (this.props.x[i] % 1 === 0) {
                context.lineWidth = 1;
                context.strokeStyle = "#ccc";
                context.beginPath();
                context.moveTo(this.props.x[i] * xScaleFactor, 0);
                context.lineTo(this.props.x[i] * xScaleFactor, this.props.height);
                context.stroke();
            }

            context.lineWidth = 2;
            context.strokeStyle = "#23c";
            context.beginPath();
            context.moveTo(this.props.x[i - 1] * xScaleFactor, halfHeight - values[i - 1] * yScaleFactor);
            context.lineTo(this.props.x[i] * xScaleFactor, halfHeight - values[i] * yScaleFactor);
            context.stroke();
        }
    }

    render() {
        return <canvas className="LinePlotArea" ref={this.updateCanvasNode} style={{width: this.props.width, height: this.props.height}} />;
    }
}

LinePlotArea.displayName = "MeshPlotArea";

LinePlotArea.propTypes = {
    height: React.PropTypes.number,
    width: React.PropTypes.number,
    x: React.PropTypes.arrayOf(React.PropTypes.number),
    y: React.PropTypes.arrayOf(React.PropTypes.number)
};

LinePlotArea.defaultProps = {
    height: 300,
    width: 800,
    x: [],
    y: []
};

module.exports = LinePlotArea;
