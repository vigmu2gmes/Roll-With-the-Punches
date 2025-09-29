-- Put functions in this file to use them in several other scripts.
-- To get access to the functions, you need to put:
-- require "my_directory.my_file"
-- in any script using the functions.

local game = {}

function init(self)
	factory.create("#face_picker", vmath.vector3(-300, -300, 0))

	factory.create("#roll_circle", vmath.vector3(640, 360, 0), nil, {}, .20)
	factory.create("#player_factory", vmath.vector3(300, 360, 0.1), nil, {}, .5)

	math.randomseed(os.time())
end


function on_message(self, message_id, message, sender)
	if message_id == hash("spawn_enemies") then
		game.enemy_count = message.spawning
		for x = 1, message.spawning do
			self.enemy = factory.create("#enemy_factory", vmath.vector3(math.random(225, 1075), math.random(100, 500), 0.1))
			--msg.post(msg.url("main", self.enemy, "enemy_script"), "enemy_count", { count = message.spawning })
		end
	end
end

return game