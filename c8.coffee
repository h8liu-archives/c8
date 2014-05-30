redraw = (timestamp) ->
    winh = $(window).height()
    console.log(winh)
    $("canvas#console").height(winh - 50)
    window.requestAnimationFrame(redraw)
    return

redraw()
window.requestAnimationFrame(redraw)
