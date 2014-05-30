c8 = new Object()

main = ->
    c8.term = new Terminal($("canvas#console")[0])
    c8.console = new Console(c8.term)
    redraw()
    return

Terminal = (canvas) ->
    self = this

    dpr = window.devicePixelRatio
    context = canvas.getContext('2d')
    context.scale(dpr, dpr)

    fontSize = 14 * dpr
    charHeight = fontSize
    lineMargin = 3 * dpr
    lineHeight = charHeight + lineMargin
    context.font = '' + fontSize + 'px Consolas'
    charWidth = context.measureText('M').width # calculate char width

    width = 0
    height = 0
    nrow = 0
    ncol = 0

    self.sayHello = ->
        s = "j$ Hello world!"
        i = 0
        for c in s.split('')
            self.drawChar(0, i, c)
            self.drawChar(1, i, c)
            i++
        return
    
    self.updateSize = (w, h) ->
        if (w == width && h == height)
            return false

        canvas.width = w * dpr
        canvas.height = h * dpr
        context.font = '' + fontSize + 'px Consolas'
        context.fillStyle = '#ddd'
        context.textBaseline = 'bottom'

        width = w
        height = h
        nrow = Math.floor(height / lineHeight)
        ncol = Math.floor(width / charWidth)
        return true
    
    self.inRange = (row, col) ->
        if row < 0
            return false
        if row >= nrow
            return false
        if col < 0
            return false
        if col >= ncol
            return false
        return true

    self.charPos = (row, col) ->
        x = col * charWidth
        y = row * lineHeight + dpr
        return { x:x, y:y }

    self.clearChar = (row, col, c) ->
        if !self.inRange(row, col)
            return
        p = self.charPos(row, col)
        context.clearRect(p.x, p.y, charWidth, charHeight)
        return

    self.drawChar = (row, col, c) ->
        if !self.inRange(row, col)
            return
        p = self.charPos(row, col)
        context.clearRect(p.x, p.y, charWidth, charHeight)
        context.fillStyle = '#ddd'
        context.fillText(c, p.x, p.y + charHeight)
        return

    self.drawCursor = (row, col, c) ->
        if !self.inRange(row, col)
            return
        p = self.charPos(row, col)
        context.fillRect(p.x, p.y + charHeight,
            charWidth, dpr)
        return
    
    self.clearCursor = (row, col, c) ->
        if !self.inRange(row, col)
            return
        p = self.charPos(row, col)
        context.clearRect(p.x, p.y + charHeight,
            charWidth, dpr)
        return

    self.nrow = -> nrow
    self.ncol = -> ncol

    return

TermBuffer = (nrow, ncol) ->
    self = this
    
    self.resize = (nrow, ncol) ->
        self.nrow = nrow
        self.ncol = ncol
        self.buf = []
        for i in [1..nrow]
            line = []
            for j in [1..ncol]
                line.push(' ')
            self.buf.push(line)
        return

    self.setChar = (row, col, c) ->
        self.buf[row][col] = c
        return

    self.drawAll = (term) ->
        for i in [0..self.nrow-1]
            for j in [0..self.ncol-1]
                term.drawChar(i, j, self.buf[i][j])
        return

    self.drawChar = (term, row, col, c) ->
        self.setChar(row, col, c)
        term.drawChar(row, col, c)
        return

    # init
    self.resize(nrow, ncol)

    return

# a rolling console line buffer
Console = (term) ->
    self = this
    self.term = term
    self.maxLines = 100000
    self.lines = []
    self.nrow = 0
    self.ncol = 0
    self.curLine = ''
    self.curPos = 0
    self.curShowing = false

    nline = (s) -> Math.ceil(s.length / ncol)

    self.printLine = (row, s) ->
        chars = s.split('')
        col = 0
        ncol = self.term.ncol()
        for c in chars
            self.term.drawChar(row, col, c)
            col++
            if col == ncol
                row++
                col = 0
        return

    self.insertChar = (c) ->
        n = self.curLine.length
        if self.curPos == n
            self.curLine = self.curLine + c
            self.curPos++
        else
            before = self.curLine.slice(0, self.curPos)
            after = self.curLine.slice(self.curPos, n)
            self.curLine = before + c + after
            self.curPos++
        return

    self.moveLeft = ->
        if self.curPos > 0
            self.curPos--
        return

    self.moveRight = ->
        n = self.curLine.length
        if self.curPos < n
            self.curPos++
        return
    
    self.newLine = ->
        self.lines.push(self.curline)
        if self.lines.length > self.maxLines
            self.lines.shift()
        self.curline = ''
        return

    self.redraw = ->
        t = ms()
        if t < 500
            if !self.curShowing
                self.curShowing = true
                term.drawCursor(0, 0)
        else
            if self.curShowing
                self.curShowing = false
                term.clearCursor(0, 0)
        return

    self.onResize = (nrow, ncol) ->
        self.nrow = nrow
        self.ncol = ncol
        self.redraw()
        return

    return

ms = ->
    d = new Date()
    return d.getMilliseconds()

redraw = (timestamp) ->
    winh = $(window).height()
    c = $("canvas#console")
    c.height(winh - 50)

    term = c8.term
    cons = c8.console

    if term.updateSize(c.width(), c.height())
        nrow = term.nrow()
        ncol = term.ncol()
        cons.onResize(nrow, ncol)
        term.sayHello()
    cons.redraw()

    window.requestAnimationFrame(redraw)
    return

$(document).ready(main)
