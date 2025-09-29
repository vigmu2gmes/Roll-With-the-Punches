components {
  id: "player"
  component: "/zAssets/player/assets/player.script"
}
components {
  id: "player_gui"
  component: "/zAssets/player/assets/gui/player_gui.gui"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"player\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/zAssets/player/assets/player.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.1
    y: 0.1
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"player\"\n"
  "mask: \"roll\"\n"
  "mask: \"enemy\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 45.0\n"
  "  data: 60.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
