local player = {}
local anim = require("library.anim8")
local love = require("love")

love.graphics.setDefaultFilter("nearest", "nearest")
player.x = 300
player.y = 200
player.speed = 5
player.spritesheet = love.graphics.newImage("assets/images/player.png")
player.grid = anim.newGrid(16, 16, player.spritesheet:getWidth(), player.spritesheet:getHeight())

player.animations = {}
player.animations.down = anim.newAnimation(player.grid('1-4', 4), 0.2)
player.animations.right = anim.newAnimation(player.grid('1-4', 1), 0.2)
return player