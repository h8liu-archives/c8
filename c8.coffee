ani = (timestamp) ->
    winh = $(window).height()
    console.log(winh)
    $("canvas#console").height(winh - 50)
    window.requestAnimationFrame(ani)
    return

ani()
window.requestAnimationFrame(ani)
