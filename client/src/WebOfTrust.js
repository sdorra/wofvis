import React, { Component } from "react";
import Graph from "react-graph-vis";
import * as gravatar from "gravatar";

class WebOfTrust extends Component {
    constructor() {
        super();
        this.state = {
            nodes: []
        };
    }

    componentDidMount() {
        fetch("nodes.json")
            .then(response => response.json())
            .then(nodes =>
                this.setState({
                    nodes: nodes
                })
            );
    }

    createGraph(nodes) {
        const gravatarOpts = {
            size: "48",
            d: "identicon"
        };

        const edges = [];
        for (let node of nodes) {
            node.label = node.name;
            node.shape = "circularImage";
            node.image = gravatar.url(node.email, gravatarOpts);
            if (node.signedBy) {
                for (let sign of node.signedBy) {
                    edges.push({
                        from: sign,
                        to: node.id
                    });
                }
            }
        }

        return {
            nodes,
            edges
        };
    }

    render() {
        const { nodes } = this.state;
        if (!nodes) {
            return <div />;
        }

        const graph = this.createGraph(nodes)

        const options = {
            layout: {
                hierarchical: false
            },
            edges: {
                color: {
                    color:'#23A3DD',
                    highlight: '#00426B'
                }
            },
            nodes: {
                font: {
                    color: '#00426B'
                }
            },
            interaction: {
                zoomView: false
            }
        };

        return (
            <div className="box">
                <Graph graph={graph} options={options} style={{ height: "520px" }} />
            </div>
        );
    }
}

export default WebOfTrust;
