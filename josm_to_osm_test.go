package josm_to_osm_converter

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert(t *testing.T) {
	t.Run("convert bytes", func(t *testing.T) {
		osmRoot, err := Convert(bytes.NewReader(testData))

		assert.NoError(t, err, "got error.")
		assert.Equal(t, 26, len(osmRoot.Nodes), "Root has 26 nodes.")
		assert.Equal(t, 7, len(osmRoot.Ways), "Root has 7 ways.")
		assert.Equal(t, 5, len(osmRoot.Relations), "Root has 5 relations.")
	})
}

func TestConvertToSql(t *testing.T) {
	t.Run("convert bytes to sql", func(t *testing.T) {
		osmRoot, err := Convert(bytes.NewReader(testData))

		assert.NoError(t, err, "Got error.")
		query, err := ConvertToSql(osmRoot)
		assert.NoError(t, err, "Got error.")
		assert.NotEmpty(t, query, "empty.")
	})
}

var testData = []byte(`<?xml version='1.0' encoding='UTF-8'?>
<osm version='0.6' generator='JOSM'>
  <bounds minlat='35.60348' minlon='139.7278' maxlat='35.61147' maxlon='139.74207' origin='OpenStreetMap server' />
  <node id='-102235' action='modify' visible='true' lat='35.60900471849' lon='139.73512583976' />
  <node id='-102238' action='modify' visible='true' lat='35.60681665956' lon='139.73418702397' />
  <node id='-102242' action='modify' visible='true' lat='35.60516285438' lon='139.73349855905' />
  <node id='-102247' action='modify' visible='true' lat='35.60984430661' lon='139.73526666212'>
    <tag k='name' v='point 2' />
  </node>
  <node id='-102248' action='modify' visible='true' lat='35.61031498094' lon='139.73747287923' />
  <node id='-102249' action='modify' visible='true' lat='35.60948811878' lon='139.73720688142' />
  <node id='-102251' action='modify' visible='true' lat='35.60871213268' lon='139.73680006125' />
  <node id='-102254' action='modify' visible='true' lat='35.60766899195' lon='139.73623677177' />
  <node id='-102255' action='modify' visible='true' lat='35.60676577375' lon='139.73608030248' />
  <node id='-102257' action='modify' visible='true' lat='35.60600248282' lon='139.73573607002' />
  <node id='-102259' action='modify' visible='true' lat='35.60983158564' lon='139.739084513' />
  <node id='-102260' action='modify' visible='true' lat='35.60899199739' lon='139.73891239677' />
  <node id='-102262' action='modify' visible='true' lat='35.60802518787' lon='139.73836475423' />
  <node id='-102264' action='modify' visible='true' lat='35.61018777193' lon='139.74077438142'>
    <tag k='name' v='point 3' />
  </node>
  <node id='-102265' action='modify' visible='true' lat='35.6091700926' lon='139.74053967748' />
  <node id='-102266' action='modify' visible='true' lat='35.60819056401' lon='139.74033626739' />
  <node id='-102268' action='modify' visible='true' lat='35.60698203819' lon='139.73972603712' />
  <node id='-102276' action='modify' visible='true' lat='35.60909376613' lon='139.73229374546'>
    <tag k='name' v='point 1' />
  </node>
  <node id='-102277' action='modify' visible='true' lat='35.60826689135' lon='139.73180869063' />
  <node id='-102278' action='modify' visible='true' lat='35.60794886031' lon='139.73116716651' />
  <node id='-102280' action='modify' visible='true' lat='35.60727463033' lon='139.73127669502' />
  <node id='-102282' action='modify' visible='true' lat='35.60721102343' lon='139.73223115774' />
  <node id='-102284' action='modify' visible='true' lat='35.60765627064' lon='139.7326849187' />
  <node id='-102289' action='modify' visible='true' lat='35.60773259848' lon='139.72842895379' />
  <node id='-102290' action='modify' visible='true' lat='35.60634596464' lon='139.73004058756' />
  <node id='-102292' action='modify' visible='true' lat='35.60592615332' lon='139.73183998449' />
  <way id='-102239' action='modify' visible='true'>
    <nd ref='-102235' />
    <nd ref='-102238' />
    <nd ref='-102242' />
    <tag k='name' v='way 3' />
  </way>
  <way id='-102250' action='modify' visible='true'>
    <nd ref='-102248' />
    <nd ref='-102249' />
    <nd ref='-102251' />
    <tag k='name' v='way 4' />
  </way>
  <way id='-102256' action='modify' visible='true'>
    <nd ref='-102254' />
    <nd ref='-102255' />
    <nd ref='-102257' />
    <tag k='name' v='way 5' />
  </way>
  <way id='-102261' action='modify' visible='true'>
    <nd ref='-102259' />
    <nd ref='-102260' />
    <nd ref='-102262' />
    <tag k='name' v='way 6' />
  </way>
  <way id='-102267' action='modify' visible='true'>
    <nd ref='-102265' />
    <nd ref='-102266' />
    <nd ref='-102268' />
    <tag k='name' v='way 7' />
  </way>
  <way id='-102279' action='modify' visible='true'>
    <nd ref='-102277' />
    <nd ref='-102278' />
    <nd ref='-102280' />
    <nd ref='-102282' />
    <nd ref='-102284' />
    <nd ref='-102277' />
    <tag k='name' v='way 2' />
  </way>
  <way id='-102291' action='modify' visible='true'>
    <nd ref='-102289' />
    <nd ref='-102290' />
    <nd ref='-102292' />
    <tag k='name' v='way 1' />
  </way>
  <relation id='-102303' action='modify' visible='true'>
    <member type='node' ref='-102247' role='' />
    <member type='way' ref='-102239' role='' />
    <tag k='name' v='relation 1' />
  </relation>
  <relation id='-102379' action='modify' visible='true'>
    <member type='way' ref='-102250' role='' />
    <member type='way' ref='-102256' role='' />
    <tag k='name' v='relation 2' />
  </relation>
  <relation id='-102388' action='modify' visible='true'>
    <member type='way' ref='-102261' role='' />
    <tag k='name' v='relation 3' />
  </relation>
  <relation id='-102395' action='modify' visible='true'>
    <member type='way' ref='-102267' role='' />
    <tag k='name' v='relation 4' />
  </relation>
  <relation id='-102410' action='modify' visible='true'>
    <member type='relation' ref='-102379' role='' />
    <member type='relation' ref='-102303' role='' />
    <tag k='name' v='parent relation 1' />
  </relation>
</osm>
`)
