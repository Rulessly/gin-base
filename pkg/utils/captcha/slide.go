package captcha

import (
	"github.com/wenlng/go-captcha-assets/resources/imagesv2"
	"github.com/wenlng/go-captcha-assets/resources/tiles"
	"github.com/wenlng/go-captcha/v2/slide"
	"log"
)

// SlideCapt 滑动验证码
var SlideCapt slide.Captcha

func init() {
	builder := slide.NewBuilder(
	//slide.WithGenGraphNumber(2),//设置图形个数
	//slide.WithEnableGraphVerticalRandom(true)//设置图形水平方向是否随机排序,
	)

	// 背景图片
	images, err := imagesv2.GetImages()
	if err != nil {
		log.Fatalln(err)
	}

	graphs, err := tiles.GetTiles()
	if err != nil {
		log.Fatalln(err)
	}

	var newGraphs = make([]*slide.GraphImage, 0, len(graphs))
	for i := 0; i < len(graphs); i++ {
		graph := graphs[i]
		newGraphs = append(newGraphs, &slide.GraphImage{
			OverlayImage: graph.OverlayImage,
			MaskImage:    graph.MaskImage,
			ShadowImage:  graph.ShadowImage,
		})
	}

	// set resources
	builder.SetResources(
		slide.WithGraphImages(newGraphs),
		slide.WithBackgrounds(images),
	)

	SlideCapt = builder.Make()
}
