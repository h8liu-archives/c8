c9 = new Object()

main = ->
    c9.term = new Terminal($("canvas#console")[0])
    redraw()
    window.requestAnimationFrame(redraw)
    return

Terminal = (canvas) ->
    self = this

    dpr = window.devicePixelRatio
    context = canvas.getContext('2d')
    context.scale(dpr, dpr)

    fontSize = 13 * dpr
    charHeight = fontSize
    lineMargin = 2 * dpr
    lineHeight = charHeight + lineMargin
    context.font = '' + fontSize + 'px Consolas'
    charWidth = context.measureText('M').width # calculate char width

    width = 0
    height = 0

    self.sayHello = ->
        s = "$ Hello world!"
        i = 0
        for c in s.split('')
            self.drawChar(0, i, c)
            i++
        return
    
    self.updateSize = (w, h) ->
        if (w == width && h == height)
            return

        canvas.width = w * dpr
        canvas.height = h * dpr
        context.font = '' + fontSize + 'px Consolas'
        context.fillStyle = '#eee'
        context.textBaseline = 'bottom'

        weight = w
        height = h
        return

    self.drawChar = (row, col, c) ->
        x = col * charWidth
        y = row * lineHeight + dpr
        context.clearRect(x, y, charWidth, charHeight)
        context.fillText(c, x, y + charHeight)
        return

    return

redraw = (timestamp) ->
    winh = $(window).height()
    c = $("canvas#console")
    c.height(winh - 50)
    term = c9.term
    term.updateSize(c.width(), c.height())
    term.sayHello()

    window.requestAnimationFrame(redraw)
    return

$(document).ready(main)
