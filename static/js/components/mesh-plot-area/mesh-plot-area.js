"use strict";

var React = require("react");
var Three = require("three");

require("./../mesh-plot-area/mesh-plot-area.less");

class MeshPlotArea extends React.Component {
    constructor(props) {
        super(props);

        this.canvas = null;

        this.componentDidMount = this.componentDidMount.bind(this);
        this.componentDidUpdate = this.componentDidUpdate.bind(this);
        this.animate = this.animate.bind(this);

        this.updateCanvasNode = node => this.canvas = node;

        this.scene = new Three.Scene();
        this.camera = new Three.OrthographicCamera(window.innerWidth / -8, window.innerWidth / 8, window.innerHeight / 8, window.innerHeight / -8, -10000, 10000);
        this.textureLoader = new Three.TextureLoader();

        var directionalLight1 = new Three.DirectionalLight(0xffffff, 0.5);
        directionalLight1.position.set(0, 1, 1);
        this.scene.add(directionalLight1);
        var pointLight = new Three.PointLight(0xffffff, 0.9, 170, 1);
        pointLight.position.set(20, 50, 100);
        this.scene.add(pointLight);
    }

    animate() {
        requestAnimationFrame(this.animate);

        // Draw the mesh
        this.renderer.render(this.scene, this.camera);
    }

    componentDidMount() {
        this.componentDidUpdate();
        this.animate(); // start the loop
    }

    componentDidUpdate() {
        // Get the three renderer ready to render our mesh
        this.renderer = new Three.WebGLRenderer({canvas: this.canvas});
        this.renderer.setClearColor(0xdddddd, 1);

        // Make sure the canvas properties are updated properly when the dom element changes size
        if (this.canvas.width !== this.canvas.clientWidth || this.canvas.height !== this.canvas.clientHeight) {
            this.canvas.width = this.canvas.clientWidth;
            this.canvas.height = this.canvas.clientHeight;
        }
        this.renderer.setSize(this.canvas.clientWidth, this.canvas.clientHeight);

        // Find the largest absolute value for scaling. We scale by the furthest value from the "ground" plane.
        var maxAbs = this.props.values[0];
        for (var i = 0; i < this.props.values.length; i++) {
            if (Math.abs(this.props.values[i]) > maxAbs) {
                maxAbs = Math.abs(this.props.values[i]);
            }
        }

        // Calculate the mesh from the parameters
        var aspectRatio = this.props.numx / this.props.numy;
        var geometry = new Three.PlaneBufferGeometry(100 * aspectRatio, 100, this.props.numx - 1, this.props.numy - 1);
        var vertices = geometry.attributes.position.array;
        for (i = 0; i < this.props.values.length; i++) {
            vertices[i * 3 + 2] = (this.props.values[i] / maxAbs) * 10;
        }
        geometry.rotateZ(-Math.PI / 4);
        geometry.rotateX(-Math.PI / 4);

        // Update the scene with the newly created mesh
        this.scene.remove(this.scene.getObjectByName("terrain"));
        var mesh = new Three.Mesh(geometry, new Three.MeshPhongMaterial({map: this.textureLoader.load(require("../../../img/textures/toonGrass.jpg"))}));
        mesh.name = "terrain";
        this.scene.add(mesh);
    }

    render() {
        return <canvas className="MeshPlotArea" ref={this.updateCanvasNode} style={{width: this.props.width, height: this.props.height}} />;
    }
}

MeshPlotArea.displayName = "MeshPlotArea";

MeshPlotArea.propTypes = {
    height: React.PropTypes.number,
    numx: React.PropTypes.number,
    numy: React.PropTypes.number,
    values: React.PropTypes.arrayOf(React.PropTypes.number),
    width: React.PropTypes.number
};

MeshPlotArea.defaultProps = {
    height: 600,
    numx: 0,
    numy: 0,
    width: 800,
    values: []
};

module.exports = MeshPlotArea;
