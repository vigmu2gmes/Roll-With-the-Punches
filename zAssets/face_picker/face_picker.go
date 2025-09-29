components {
  id: "face_picker"
  component: "/zAssets/face_picker/face_picker.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"face_picker\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/zAssets/ground/environment.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.27
    y: 0.27
  }
}
