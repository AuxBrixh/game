local love = require("love")
local player = require("player")


function love.load()
    local camera = require("library.camera")
    cam = camera()

    background = love.graphics.newImage("assets/images/waves.jpg")
end

function love.update(dt) 
    if love.keyboard.isScancodeDown("w") then
        player.y = player.y - player.speed
        if player.y <= -2300 then
            player.y = 600
        end
    end
    if love.keyboard.isScancodeDown("s") then
        player.y = player.y + player.speed
        if player.y >= 2300 then
            player.y = -600
        end
    end
    if love.keyboard.isScancodeDown("a") then
        player.x = player.x - player.speed
        if player.x <= -2300 then
            player.x = 600
        end
    end
    if love.keyboard.isScancodeDown("d") then
        player.x = player.x + player.speed
        if player.x >= 2300 then
            player.x = -600
        end
    end

    player.animations.right:update(dt)
    cam:lookAt(player.x, player.y)
end

function love.draw()
    cam:attach()
        love.graphics.draw(background, 0, 0)
        player.animations.right:draw(player.spritesheet, player.x, player.y, nil, 10)
    cam:detach()
end
