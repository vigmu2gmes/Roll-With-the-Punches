components {
  id: "enemy_script"
  component: "/zAssets/enemies/assets/enemy_script.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"enemy\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/zAssets/enemies/assets/enemies.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.05
    y: 0.05
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"enemy\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 17.5\n"
  "  data: 30.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
