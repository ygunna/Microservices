import { Injectable, EventEmitter } from '@angular/core';
import * as d3 from 'd3';
import { environment } from '../../../environments/environment';

let svg, zoom;
let currentNode;
let dragLineId: number = 0;
const rectWidth: number = 150;
const rectHeight: number = 25;
let droppedSvgX = 0, droppedSvgY = 0;
let simulation;

const FORCES = {
  LINKS: 1 / 50,
  COLLISION: 1,
  CHARGE: -100
};

@Injectable()
export class D3Service {

  public static selectedElement: {type: string, element: any};

  constructor() { }

  applyZoomableBehaviour(svgElement, containerElement) {
    let container, zoomed;

    svg = d3.select(svgElement);
    container = d3.select(containerElement);

    zoomed = () => {
      const transform = d3.event.transform;
      container.attr('transform', 'translate(' + transform.x + ',' + transform.y + ') scale(' + transform.k + ')');
    };

    zoom = d3.zoom().scaleExtent([1 / 5, 5]).on('zoom', zoomed);
    svg.call(zoom);

  }
  zoomClick(direction) {
    direction = (direction === 'in') ? 1 : -1;
    zoom.scaleBy(svg, 1+(direction/10));
  }

  nodeSimulation(width, height, nodes, links) {
    simulation = d3.forceSimulation()
      .force('collide',d3.forceCollide().strength(FORCES.COLLISION).radius(d => d['r'] + 5).iterations(2))
      .force('charge', d3.forceManyBody().strength(d => FORCES.CHARGE * d['r']).distanceMin(10).distanceMax(1000))
      .force('link', d3.forceLink().id(function(d) { return d['id']; }).strength(FORCES.LINKS))
      .force('center', d3.forceCenter(width / 2, height / 2))
      .force('x', d3.forceX(width / 2))
      .force('y', d3.forceY(height / 2))
      .on('tick', ticked)
      .on('end', tickend);

    var link = svg.selectAll(".link-path"),
      node = svg.selectAll(".node");

    simulation.nodes(nodes);
    simulation.force("link").links(links);

    link = link.data(links);
    node = node.data(nodes);
    function ticked() {
      link.attr('d', function(d) { return 'M'+d.source.x+','+d.source.y+'L'+d.target.x+','+d.target.y; });
    }
    function tickend() {
      let map = new Map();
      if(link['_groups'].length > 0) {
        let groups = link['_groups'][0];
        for(let group of groups) {
          let data = group['__data__'];
          map.set(data.id, data);
        }
      }
      for(let link of links) {
        link['sNode']['x'] = map.get(link['id'])['source']['x'];
        link['sNode']['y'] = map.get(link['id'])['source']['y'];
        link['tNode']['x'] = map.get(link['id'])['target']['x'];
        link['tNode']['y'] = map.get(link['id'])['target']['y'];
      }
    }
  }

  applyDraggableBehaviour(element, node, nodes, links, editComponent) {
    let gapX, gapY, dragLine;
    const d3element = d3.select(element);
    const parent = d3element.select(function() { return this.parentNode; });
    const grandParent = parent.select(function() { return this.parentNode; });

    d3.select("svg").selectAll("circle,rect").on("mouseover", function(d){
      currentNode = d3.select(this);
    });

    function started() {
      D3Service.selectedElement = {type: 'node', element: node};
      d3.selectAll('circle, rect').classed("active", false);
      d3.selectAll('#info').classed("visible", false);
      d3.selectAll('#close').classed("visible", false);
      d3.selectAll('path').classed("selected", false);

      if(d3.event.sourceEvent.shiftKey) {
        simulation.stop();
        dragLine = grandParent.append('g').attr('linkPath', '').attr('id','tteesstt').append('svg:path');
        dragLine.attr('style', 'stroke: #333; stroke-width: 3px; marker-end: url(#mark-end-arrow);').attr('d', 'M0,0L0,0');
        d3.event.on('drag', lineDragged).on('end', lineEnded);
      } else {
        d3.event.on('drag', dragged).on('end', ended);

        gapX = node.x - d3.event.x;
        gapY = node.y - d3.event.y;
      }
      parent.raise();
      d3element.select(node.shape).classed("active", true);
      d3element.select('#info').classed("visible", true);
      d3element.select('#close').classed("visible", true);

      function dragged() {
        d3element.attr("cx", node.x = d3.event.x + gapX).attr("cy", node.y = d3.event.y + gapY);
        updateGraph(node);
      }
      function ended() {
        gapX = 0;
        gapY = 0;
      }

      function lineDragged() {
        var optX = 0, optY = 0;
        if(node.shape == 'rect') {
          optX = (rectWidth / 2), optY = (rectHeight / 2);
        }
        setLine(dragLine, node.x + optX, node.y + optY, d3.mouse(this)[0], d3.mouse(this)[1]);
      }
      function lineEnded() {
        dragLineId += 1;
        const parentDragLine = dragLine.select(function() { return this.parentNode; });
        parentDragLine.remove();
        if(currentNode == null) {
          return false;
        }

        var sourceId = d3element.attr("id");
        var targetG = currentNode.select(function() { return this.parentNode; });
        var targetId = targetG.select(function() { return this.parentNode; }).attr("id");
        if(sourceId != targetId) {
          var sOptX = 0, sOptY = 0, tOptX = 0, tOptY = 0;
          if(node.shape == 'rect') {
            sOptX = (rectWidth / 2), sOptY = (rectHeight / 2);
          }
          if(currentNode.node().nodeName == 'rect') {
            tOptX = (rectWidth / 2), tOptY = (rectHeight / 2);
          }
          var link = {
            id: dragLineId,
            sNode: {
              id: sourceId,
              x: Number(d3element.select('g').attr('x')) + sOptX,
              y: Number(d3element.select('g').attr('y')) + sOptY
            },
            tNode: {
              id: targetId,
              x: Number(targetG.attr('x')) + tOptX,
              y: Number(targetG.attr('y')) + tOptY
            }
          };
          for(var _link of links) {
            if(_link.sNode.id == sourceId && _link.tNode.id == targetId) {
              return false;
            }
          }
          links.push(link);
          editComponent.putLink(link);

          var targetType;
          for(var n of nodes) {
            if(targetId == n.id) {
              targetType = n.type;
              break;
            }
          }
          setTimeout(function() {
            if(targetType == environment.nodeTypeService) {
              d3.select('#link-path-'+link.id).classed('service', true);
            } else {
              d3.select('#link-path-'+link.id).classed('app', true);
            }
          });
          if(currentNode.node().nodeName == 'rect') {
            var incidenceAngle = getIncidenceAngle(link);
            var refX = getRefXtoEdge(incidenceAngle);
            var defs = d3.select("svg").select('defs');
            defs.append('svg:marker')
              .attr('id', 'rect-arrow-'+link.id)
              .attr('viewBox', '0 -5 10 10')
              .attr('refX', refX)
              .attr('markerWidth', 3.5)
              .attr('markerHeight', 3.5)
              .attr('orient', 'auto')
              .append('svg:path')
              .attr('d', 'M0,-5L10,0L0,5');

            setTimeout(function() {
              var path = d3.select('#link-'+link.id).select('path');
              path.style('marker-end', 'url(#rect-arrow-'+link.id+')');
            });
          }

        }
      }
    }

    d3element.call(d3.drag().on('start', started));

    function setLine(line, nodeX, nodeY, arrowX, arrowY) {
      line.attr('d', 'M' + nodeX + ',' + nodeY + 'L' + arrowX + ',' + arrowY);
    }
    function updateGraph(node) {
      var optX = 0, optY = 0;
      if(node.shape == 'rect') {
        optX = (rectWidth / 2), optY = (rectHeight / 2);
      }
      for(var link of links) {
        if(link.sNode.id == node.id) {
          var _path = d3.select('#link-'+link.id).select('path');
          var _d = _path.attr('d');
          var _l = _d.substring(_d.indexOf('L'));
          _path.attr('d', 'M' + (node.x + optX) + ',' + (node.y + optY) + _l);
          link.sNode.x = (node.x + optX);
          link.sNode.y = (node.y + optY);
        }
        if(link.tNode.id == node.id) {
          var _path = d3.select('#link-'+link.id).select('path');
          var _d = _path.attr('d');
          var _m = _d.substring(0,_d.indexOf('L'));
          _path.attr('d', _m + 'L' + (node.x + optX) + ',' + (node.y + optY));
          link.tNode.x = (node.x + optX);
          link.tNode.y = (node.y + optY);
        }

        var incidenceAngle = getIncidenceAngle(link);
        var refX = getRefXtoEdge(incidenceAngle);
        d3.select('#rect-arrow-'+link.id).attr('refX', refX);
      }
    }

    function getIncidenceAngle(link) {
      var dy = link.tNode.y - link.sNode.y;
      var dx = link.tNode.x - link.sNode.x;
      var angle = Math.atan(dy / dx) * (180.0 / Math.PI);

      if(dx < 0.0) {
        angle += 180.0;
      } else {
        if(dy < 0.0) {
          angle += 360.0;
        }
      }
      return angle;
    }
    function getRefXtoEdge(incidenceAngle) {
      var refX;

      // 연결선이 사각형 노드의 모서리에 닿을 때 각도. 밑변과 높이의 정의 기준이 됨
      var edgeAngle = Math.atan(rectHeight / rectWidth) * 180/Math.PI;

      switch(true) {
        case incidenceAngle < edgeAngle: {
          refX = (rectWidth / 2) / Math.cos(incidenceAngle * Math.PI/180);
          break;
        }
        case edgeAngle <= incidenceAngle && incidenceAngle < 90: {
          refX = (rectHeight / 2) / Math.cos((90 - incidenceAngle) * Math.PI/180);
          break;
        }
        case 90 <= incidenceAngle && incidenceAngle < (180 - edgeAngle): {
          refX = (rectHeight / 2) / Math.cos((incidenceAngle - 90) * Math.PI/180);
          break;
        }
        case (180 - edgeAngle) <= incidenceAngle && incidenceAngle < 180: {
          refX = (rectWidth / 2) / Math.cos((180 - incidenceAngle) * Math.PI/180);
          break;
        }
        case 180 <= incidenceAngle && incidenceAngle < (180 + edgeAngle): {
          refX = (rectWidth / 2) / Math.cos((incidenceAngle - 180) * Math.PI/180);
          break;
        }
        case (180 + edgeAngle) <= incidenceAngle && incidenceAngle < 270: {
          refX = (rectHeight / 2) / Math.cos((270 - incidenceAngle) * Math.PI/180);
          break;
        }
        case 270 <= incidenceAngle && incidenceAngle < (360 - edgeAngle): {
          refX = (rectHeight / 2) / Math.cos((incidenceAngle - 270) * Math.PI/180);
          break;
        }
        case (360 - edgeAngle) <= incidenceAngle && incidenceAngle < 360: {
          refX = (rectWidth / 2) / Math.cos((360 - incidenceAngle) * Math.PI/180);
          break;
        }
      }
      return refX + 11;
    }
  }

  applyAppDroppableBehaviour(event, jsonObject, droppedNodes, msaName) {
    simulation.stop();
    d3.selectAll('circle,rect').classed("active", false);
    d3.selectAll('#info').classed("visible", false);
    d3.selectAll('#close').classed("visible", false);

    let _svg = document.querySelector('svg');
    svg.on('mousemove', function() {
      var _pt = _svg.createSVGPoint();
      _pt.x = d3.event.clientX;
      _pt.y = d3.event.clientY;
      var loc = _pt.matrixTransform((document.querySelector('svg > g') as SVGGraphicsElement).getScreenCTM().inverse());
      droppedSvgX = loc.x;
      droppedSvgY = loc.y;
    });

    setTimeout(function() {
      for(var node of droppedNodes) {
        if(node.id === jsonObject['id'] || node.id === 'INITIAL_'+jsonObject['id']) {
          selectNode(jsonObject);
          return false;
        }
      }
      var name = '';
      if(jsonObject['type'] == environment.nodeTypeService) {
        name = jsonObject['label'];
      } else {
        name = jsonObject['name'];
      }

      // duplication check
      for(var node of droppedNodes) {
        if(node.name == name + '-' + msaName){
          return false;
        }
      }

      droppedNodes.push({
        shape: jsonObject['shape'],
        x: droppedSvgX, //event.offsetX,
        y: droppedSvgY, //event.offsetY,
        r: 25,
        id: 'INITIAL_'+jsonObject['id'],
        name: name + '-' + msaName,
        type: jsonObject['type'],
        color: jsonObject['color']
      });
      setTimeout(() => {
        selectNode(jsonObject);
      });
    }, 100);

    function selectNode(jsonObject) {
      var g = d3.select('svg').select("[id='INITIAL_"+jsonObject['id']+"']");
      var parent = g.select(function() { return (<HTMLElement>this).parentElement; });
      parent.raise();
      g.select(jsonObject['shape']).classed("active", true);
      g.select('#info').classed("visible", true);
      g.select('#close').classed("visible", true);
    }
  }

  applyMouseDownBehaviour() {
    d3.selectAll('circle,rect').classed("active", false);
    d3.selectAll('#info').classed("visible", false);
    d3.selectAll('#close').classed("visible", false);
    D3Service.selectedElement = null;
    d3.selectAll('path').classed("selected", false);
  }

  clickPath(path, link) {
    D3Service.selectedElement = {type: 'line', element: link};
    d3.selectAll('path').classed("selected", false);
    d3.select(path).select('path').classed('selected', true);
  }

}
